import { writable } from "svelte/store";
import * as types from "./types/tasks";

export const isModalOpen = writable(false);

export async function addTaskDatabase(details: types.TaskInformation) {
    console.log("I have been called")
    const fetchParams = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(details)
    }

    fetch("/api/tasks/new", fetchParams)
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(err => console.error(`there was an error trying to add the task: ${err}`))
}
