"use client";

import { Poll } from "@/types/poll";
import PollCard from "./PollCard";
import { useState } from "react";
import Modal from "../Modal";
import Button from "../Button";
import { Heart, Play } from "lucide-react";

export default function PollsGrid({ polls }: { polls: Poll[] }) {
    const [open, setOpen] = useState(false);
    const [selected, setSelected] = useState<Poll | null>(null);

    return (
        <div className="grid grid-cols-2 gap-3">
            {polls.map((poll) => (
                <PollCard
                    key={poll.id}
                    poll={poll}
                    onClick={() => {
                        setSelected(poll);
                        setOpen(true);
                    }}
                />
            ))}

            <Modal
                open={open}
                onClose={() => setOpen(false)}
                title={selected?.title}
                inertRootId="page-root"
            >
                {selected && (
                    <div className="flex flex-1 flex-col justify-between">
                        <div>
                            <p className="text-lg">
                                Author: {selected.author.email}
                            </p>

                            <div className="py-2">
                                <Button
                                    onClick={() => {}}
                                    className="gap-3 px-4"
                                    aria-label="Like poll"
                                >
                                    <Heart className="h-4 w-4" />
                                    <span className="">{selected.likes}</span>
                                </Button>
                            </div>
                        </div>

                        <div className="flex flex-col space-y-4">
                            <div className="space-y-2">
                                <p className="">
                                    Created:{" "}
                                    {new Date(
                                        selected.createdAt
                                    ).toLocaleString()}
                                </p>
                                <p className="">
                                    Updated:{" "}
                                    {new Date(
                                        selected.updatedAt
                                    ).toLocaleString()}
                                </p>
                            </div>
                            <Button
                                href={{
                                    pathname: `/polls/${selected.id}`,
                                }}
                            >
                                <Play className="h-4 w-4" />
                                <span className="">Participate</span>
                            </Button>
                        </div>
                    </div>
                )}
            </Modal>
        </div>
    );
}
