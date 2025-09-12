import { Poll } from "@/types/poll";

export default function PollDetails({ poll }: { poll: Poll }) {
    return (
        <div>
            <h1>{poll.title}</h1>
            <p>Author: {poll.author.email}</p>
            <p>Created at: {poll.createdAt}</p>
            <p>Updated at: {poll.updatedAt}</p>
        </div>
    );
}
