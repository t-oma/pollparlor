import { Poll } from "@/types/poll";
import PollCard from "./PollCard";

export default function PollsGrid({ polls }: { polls: Poll[] }) {
    return (
        <div className="grid grid-cols-2 gap-3">
            {polls.map((poll) => (
                <PollCard
                    key={poll.id}
                    poll={poll}
                />
            ))}
        </div>
    );
}
