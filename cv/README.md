# CV Materials

This directory contains Wahidyan Kresna Fridayoka's career evidence and professional-profile drafts. It is the source for CV- or LinkedIn-related work; read the relevant file before making claims, edits, or exports.

## Sources and Derivatives

- `cv-raw.md` is the evidence base. Preserve its dates, metrics, source wording, context, and notes about material that is not suitable for public use.
- `cv-linkedin.md` is the long-form, copy-ready LinkedIn profile.
- `cv-ats.md` is the concise ATS CV source; `cv-ats.pdf` is its two-page export.
- `generate-cv-ats-pdf.py` generates the PDF from `cv-ats.md`.
- `linkedin-projects.md` contains draft LinkedIn project entries and publishing notes.

## Working Rules

Treat `cv-raw.md` as the factual source of truth. Keep public-facing CV and LinkedIn content accurate, contextualized, and consistent with it. Do not publish or promote sensitive notes marked for exclusion without explicit owner direction.

Regenerate `cv-ats.pdf` after changing `cv-ats.md` with:

```sh
uv run --with reportlab python cv/generate-cv-ats-pdf.py
```

Inspect the generated PDF before sharing it.

## Known Source Gap

`cv-raw.md` links to `hijra-bank-okr-2026-management-dashboard.md`, which is not currently present in this directory. Do not invent its contents; add the source or verify the affected claims with the owner before relying on it.
