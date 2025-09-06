import type { PollItem } from "@/types/poll";

export default function PollMember({
    onSelect,
    item,
}: Readonly<{
    onSelect: () => void;
    item: PollItem | null;
}>) {
    return (
        <div className="flex flex-1/2 border dark:border-zinc-500">
            <button
                type="button"
                className="flex-1 cursor-pointer p-4 transition-colors dark:bg-zinc-900 dark:hover:bg-zinc-800"
                onClick={onSelect}
            >
                <span className="text-3xl">{item ? item.title : "-"}</span>
            </button>
        </div>
    );
}
