import { writable } from "svelte/store";
import * as types from "../types/tasks";


export const isModalOpen = writable(false);

export async function addTaskDatabase(details: types.TaskInformation) {
    try {
        const fetchParams = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(details)
        }

        const response = await fetch("/api/tasks", fetchParams)
        const json = response.json();
        console.log(json);
    } catch (error) {
        console.error(error)
    }
}

export async function removeTask(id: number) {
    try {
        const response = await fetch(`/api/tasks/${id}`, { method: "DELETE" })
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }
        await response.json();
    } catch (error) {
        console.error(error)
    }
}


export async function getAllTasks() {
    try {
        const response = await fetch('/api/tasks');
        const data = await response.json();
        console.log(data);
        return new Promise((fulfil) => {
            fulfil(data)
        })
    } catch (err) {
        console.error(`There was an error trying to get tasks. ${err}`);
    }
}