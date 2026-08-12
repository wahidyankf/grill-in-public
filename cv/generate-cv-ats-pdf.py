#!/usr/bin/env python3
"""Generate an ATS-safe PDF from cv-ats.md.

Usage:
    uv run --with reportlab python generate-cv-ats-pdf.py

The resulting PDF deliberately uses a single text column, standard fonts, and no tables,
graphics, or text boxes so both humans and applicant tracking systems can read it reliably.
"""

from __future__ import annotations

import re
from pathlib import Path

from reportlab.lib import colors
from reportlab.lib.enums import TA_CENTER, TA_LEFT
from reportlab.lib.pagesizes import A4
from reportlab.lib.styles import ParagraphStyle, getSampleStyleSheet
from reportlab.lib.units import mm
from reportlab.platypus import Paragraph, SimpleDocTemplate
from xml.sax.saxutils import escape


BASE_DIR = Path(__file__).resolve().parent
INPUT_PATH = BASE_DIR / "cv-ats.md"
OUTPUT_PATH = BASE_DIR / "cv-ats.pdf"


def normalize_text(value: str) -> str:
    """Make Markdown text safe for reportlab and friendlier to ATS text extraction."""
    # Escape only after recognising links. Escaping first would hide Markdown
    # delimiters, while inserting ReportLab <link> tags before escaping would
    # turn those tags into visible text.
    links: list[tuple[str, str]] = []

    def stash_link(target: str, visible_text: str | None = None) -> str:
        # Temporary markers preserve URLs through the plain-text cleanup below;
        # they are replaced by ReportLab markup only after escaping is complete.
        marker = f"__CV_URL_{len(links)}__"
        links.append((target, visible_text or target))
        return marker

    def replace_markdown_link(match: re.Match[str]) -> str:
        label, target = match.group(1), match.group(2)
        # Email labels are already readable addresses, so avoid duplicating them
        # as "address: mailto:address" in extracted ATS text.
        if target == f"mailto:{label}":
            return stash_link(target, label)
        return f"{label}: {stash_link(target)}"

    # Process Markdown links before bare URLs so a URL inside [label](URL) is
    # not recognised twice. The controlled CV source needs only this subset.
    value = re.sub(
        r"\[([^\]]+)]\(((?:https?|mailto):[^)]+)\)", replace_markdown_link, value
    )
    value = re.sub(
        r"<(https?://[^>]+)>", lambda match: stash_link(match.group(1)), value
    )
    value = re.sub(
        r"(?<![\w/])(https?://[^\s<>]+)",
        lambda match: stash_link(match.group(1)),
        value,
    )
    value = re.sub(
        r"(?<![\w@])([A-Za-z0-9.!#$%&'*+/=?^_`{|}~-]+@[A-Za-z0-9-]+(?:\.[A-Za-z0-9-]+)+)",
        lambda match: stash_link(f"mailto:{match.group(1)}", match.group(1)),
        value,
    )
    # ATS readers need text, not Markdown emphasis or typography that can map
    # inconsistently across fonts and extraction tools.
    value = re.sub(r"\*\*(.*?)\*\*", r"\1", value)
    value = value.replace("—", " - ").replace("–", "-").replace("‑", "-")
    value = escape(value)

    for index, (target, visible_text) in enumerate(links):
        # Escape URL attributes separately: XML text escaping does not protect a
        # quoted attribute if its target contains a quotation mark.
        marker = f"__CV_URL_{index}__"
        safe_target = escape(target, {'"': "&quot;"})
        safe_visible_text = escape(visible_text)
        value = value.replace(
            marker,
            f'<link href="{safe_target}" color="#1F4E79">{safe_visible_text}</link>',
        )

    return value


def add_page_number(canvas, document) -> None:
    """Add an unobtrusive, selectable footer without introducing layout elements."""
    # saveState/restoreState prevents the footer's font and colour from leaking
    # into ReportLab's normal story rendering on later pages.
    canvas.saveState()
    footer = f"Wahidyan Kresna Fridayoka | Page {document.page}"
    canvas.setFont("Helvetica", 8)
    canvas.setFillColor(colors.HexColor("#6B7280"))
    canvas.drawRightString(A4[0] - 16 * mm, 12 * mm, footer)
    canvas.restoreState()


def styles() -> dict[str, ParagraphStyle]:
    # Start from ReportLab defaults, then centralise every intentional layout
    # choice so the Markdown-to-PDF mapping below stays focused on structure.
    base = getSampleStyleSheet()
    return {
        "name": ParagraphStyle(
            "Name",
            parent=base["Normal"],
            fontName="Helvetica-Bold",
            fontSize=18,
            leading=22,
            alignment=TA_CENTER,
            spaceAfter=2,
        ),
        "headline": ParagraphStyle(
            "Headline",
            parent=base["Normal"],
            fontName="Helvetica-Bold",
            fontSize=10.5,
            leading=14,
            alignment=TA_CENTER,
            textColor=colors.HexColor("#1F4E79"),
            spaceAfter=2,
        ),
        "contact": ParagraphStyle(
            "Contact",
            parent=base["Normal"],
            fontName="Helvetica",
            fontSize=8.5,
            leading=11,
            alignment=TA_CENTER,
            spaceAfter=10,
        ),
        "section": ParagraphStyle(
            "Section",
            parent=base["Normal"],
            fontName="Helvetica-Bold",
            fontSize=11,
            leading=14,
            textColor=colors.HexColor("#1F4E79"),
            spaceBefore=9,
            spaceAfter=4,
        ),
        "employer": ParagraphStyle(
            "Employer",
            parent=base["Normal"],
            fontName="Helvetica-Bold",
            fontSize=10,
            leading=13,
            spaceBefore=5,
            spaceAfter=1,
            keepWithNext=True,
        ),
        "role": ParagraphStyle(
            "Role",
            parent=base["Normal"],
            fontName="Helvetica-Bold",
            fontSize=9.5,
            leading=12,
            spaceBefore=3,
            spaceAfter=2,
            keepWithNext=True,
        ),
        "body": ParagraphStyle(
            "Body",
            parent=base["Normal"],
            fontName="Helvetica",
            fontSize=9,
            leading=11.8,
            alignment=TA_LEFT,
            spaceAfter=4,
        ),
        "bullet": ParagraphStyle(
            "Bullet",
            parent=base["Normal"],
            fontName="Helvetica",
            fontSize=9,
            leading=11.8,
            leftIndent=10,
            firstLineIndent=-8,
            spaceAfter=2,
        ),
    }


def build_story(source: str) -> list:
    """Convert the small, controlled Markdown subset used by cv-ats.md to flowables."""
    style = styles()
    story: list = []
    paragraph_lines: list[str] = []
    plain_line_index = 0

    def flush_paragraph() -> None:
        nonlocal plain_line_index
        if not paragraph_lines:
            return
        # Markdown paragraphs can span lines; joining with spaces prevents an
        # accidental line break from becoming a word collision in the PDF.
        text = normalize_text(" ".join(paragraph_lines))
        if plain_line_index == 0:
            story.append(Paragraph(text, style["headline"]))
        elif plain_line_index == 1:
            story.append(Paragraph(text, style["contact"]))
        else:
            story.append(Paragraph(text, style["body"]))
        plain_line_index += 1
        paragraph_lines.clear()

    for raw_line in source.splitlines():
        line = raw_line.strip()
        if not line:
            flush_paragraph()
            continue

        # This is deliberately a small, controlled Markdown parser, not a
        # general renderer. Each accepted heading level maps to one CV role.
        if line.startswith("# "):
            flush_paragraph()
            story.append(Paragraph(normalize_text(line[2:]), style["name"]))
            continue
        if line.startswith("## "):
            flush_paragraph()
            story.append(Paragraph(normalize_text(line[3:]), style["section"]))
            continue
        if line.startswith("### "):
            flush_paragraph()
            story.append(Paragraph(normalize_text(line[4:]), style["employer"]))
            continue
        if line.startswith("#### "):
            flush_paragraph()
            story.append(Paragraph(normalize_text(line[5:]), style["role"]))
            continue
        if line.startswith("- "):
            flush_paragraph()
            story.append(Paragraph(f"- {normalize_text(line[2:])}", style["bullet"]))
            continue

        paragraph_lines.append(line.replace("  ", " "))

    flush_paragraph()
    return story


def main() -> None:
    # Keep the input and output beside this script so the command is portable
    # from any working directory and does not depend on the caller's CWD.
    source = INPUT_PATH.read_text(encoding="utf-8")
    document = SimpleDocTemplate(
        str(OUTPUT_PATH),
        pagesize=A4,
        rightMargin=16 * mm,
        leftMargin=16 * mm,
        topMargin=14 * mm,
        bottomMargin=18 * mm,
        title="Wahidyan Kresna Fridayoka - ATS CV",
        author="Wahidyan Kresna Fridayoka",
    )
    document.build(
        build_story(source), onFirstPage=add_page_number, onLaterPages=add_page_number
    )
    print(f"Created {OUTPUT_PATH}")


if __name__ == "__main__":
    main()
