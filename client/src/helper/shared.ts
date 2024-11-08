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

        const response = await fetch("/api/tasks/new", fetchParams)
        const json = response.json();
        console.log(json);
    } catch (error) {
        console.error(error)
    }
}

export async function removeTask(id: number) {
    try {

        const response = await fetch(`/api/task/delete/${id}`)
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }

        const json = await response.json();
        console.log(json)
    } catch (error) {
        console.error(error)
    }
}