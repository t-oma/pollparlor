"use client";

import Button from "@/components/Button";
import { useState } from "react";

export default function CreatePollForm() {
    const [itemsCount, setItemsCount] = useState(4);

    return (
        <form
            action=""
            className="flex w-full flex-1 flex-col space-y-2 p-4"
        >
            <div className="flex flex-col gap-2">
                <label htmlFor="title">Title</label>
                <input
                    type="text"
                    id="title"
                    name="title"
                    placeholder="Enter title"
                    className="border-border rounded-md border p-1"
                />
            </div>

            <div className="flex flex-col gap-2">
                {Array.from({ length: itemsCount }, (_, i) => (
                    <div
                        key={i}
                        className="flex flex-col gap-2"
                    >
                        <label htmlFor={`item${i + 1}`}>Item {i + 1}</label>
                        <input
                            type="text"
                            id={`item${i + 1}`}
                            name={`item${i + 1}`}
                            placeholder="Enter item"
                            className="border-border rounded-md border p-1"
                        />
                    </div>
                ))}
            </div>

            <div className="flex w-full flex-col gap-2 py-2">
                <Button
                    type="button"
                    onClick={() => setItemsCount((prev) => prev + 1)}
                >
                    Add item
                </Button>

                <Button type="submit">Create</Button>
            </div>
        </form>
    );
}
