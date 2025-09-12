import "server-only";

import { Poll } from "@/types/poll";
import { dateIsoToLocalString } from "@/utils/date";

const API_BASE = process.env.API_BASE_URL!;

export async function fetchPolls(): Promise<Poll[]> {
    const res = await fetch(`${API_BASE}/api/v1/polls`, {
        // cache control: cache for 60s; or use { cache: "no-store" }  to disable caching
        next: { revalidate: 60, tags: ["polls"] },
    });

    if (!res.ok) {
        console.error("Failed to fetch polls", res.status);
        return [];
    }

    const data = (await res.json()) as Poll[];
    return data.map((r) => ({
        ...r,
        createdAt: dateIsoToLocalString(r.createdAt),
        updatedAt: dateIsoToLocalString(r.updatedAt),
        author: {
            ...r.author,
            createdAt: dateIsoToLocalString(r.author.createdAt),
            passwordChangedAt: dateIsoToLocalString(r.author.passwordChangedAt),
        },
    }));
}

export async function fetchPoll(id: string): Promise<Poll | null> {
    const res = await fetch(`${API_BASE}/api/v1/polls/${id}`, {
        next: { revalidate: 60, tags: ["poll", `poll:${id}`] },
    });
    if (!res.ok) return null;
    const raw = await res.json();
    return {
        ...raw,
        createdAt: dateIsoToLocalString(raw.createdAt),
        updatedAt: dateIsoToLocalString(raw.updatedAt),
    } as Poll;
}
