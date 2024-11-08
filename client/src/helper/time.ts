import { isSameDay, isSameWeek, isSameMonth, isSameYear } from 'date-fns';
// import * as types from "../types";

export function isDay(now: Date, taskDate: Date): boolean { // this one isn't getting properly called
    //code to check what week and day of the week now is and taskDate is
    console.log("day was called")
    console.log(`current date is ${now} and the taskDate is ${taskDate}`)
    return isSameDay(now, taskDate)
}

export function isWeek(now: Date, taskDate: Date): boolean {
    console.log("week was called")
    return isSameWeek(now, taskDate)
}

export function isMonth(now: Date, taskDate: Date): boolean {
    console.log("month was called")
    return isSameMonth(now, taskDate)
}

export function isYear(now: Date, taskDate: Date): boolean {
    console.log("year was called")
    return isSameYear(now, taskDate)
}