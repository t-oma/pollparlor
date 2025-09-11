export function dateIsoToLocalString(iso: string): string {
    const d = new Date(iso);
    return d.toLocaleString("uk-UA");
}
