export function dateIsoToLocalString(iso: string): string {
    const d = new Date(iso);
    return d.toLocaleString("uk-UA");
}

export function formatRelativeDate(isoString: string): string {
    const date = new Date(isoString);
    const now = new Date();

    const diffTime = now.getTime() - date.getTime();
    const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24));
    const diffWeeks = Math.floor(diffDays / 7);
    const diffMonths = Math.floor(diffDays / 30);

    if (diffDays === 0) {
        return date.toLocaleTimeString("uk-UA", {
            hour: "2-digit",
            minute: "2-digit",
        });
    }

    if (diffDays === 1) {
        return "yesterday";
    }

    if (diffDays < 7) {
        return `${Math.abs(diffDays)} day${diffDays === -1 ? "" : "s"} ago`;
    }

    if (diffDays < 30) {
        return diffWeeks === 1 ? "last week" : `${diffWeeks} weeks ago`;
    }

    if (now.getFullYear() === date.getFullYear()) {
        return date.toLocaleDateString("uk-UA", {
            day: "numeric",
            month: "short",
        });
    }

    return date.toLocaleDateString("uk-UA", {
        day: "numeric",
        month: "short",
        year: "numeric",
    });
}

export function formatRelativeTimeSimple(isoString: string): string {
    const date = new Date(isoString);
    const now = new Date();
    const diffTime = now.getTime() - date.getTime();
    const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24));

    const rtf = new Intl.RelativeTimeFormat("uk", {
        numeric: "auto",
    });

    if (diffDays === 0) {
        return date.toLocaleTimeString("uk-UA", {
            hour: "2-digit",
            minute: "2-digit",
        });
    }

    if (diffDays < 7) {
        return rtf.format(-diffDays, "day");
    }

    if (diffDays < 30) {
        return rtf.format(-Math.floor(diffDays / 7), "week");
    }

    return date.toLocaleDateString("uk-UA", {
        day: "numeric",
        month: "short",
        year: now.getFullYear() !== date.getFullYear() ? "numeric" : undefined,
    });
}
