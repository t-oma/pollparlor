import "server-only";

import { Poll } from "@/types/poll";

const API_BASE = process.env.API_BASE_URL!;

export async function fetchPolls(): Promise<Poll[]> {
    try {
        const res = await fetch(`${API_BASE}/api/v1/polls`, {
            next: { revalidate: 60, tags: ["polls"] },
        });

        if (!res.ok) {
            console.error("Failed to fetch polls", res.status);
            return [];
        }

        return (await res.json()) as Poll[];
    } catch (error) {
        console.error("Failed to fetch polls", error);
        return [];
    }
}

export async function fetchPoll(id: string): Promise<Poll | null> {
    try {
        const res = await fetch(`${API_BASE}/api/v1/polls/${id}`, {
            next: { revalidate: 60, tags: ["poll", `poll:${id}`] },
        });
        if (!res.ok) return null;
        return (await res.json()) as Poll;
    } catch (error) {
        console.error(`Failed to fetch poll ${id}:`, error);
        return null;
    }
}
