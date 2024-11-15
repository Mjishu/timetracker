
export type TaskInformation = {
    id?: number;
    title: string;
    time: string;
    date: string;
    type: string;
}

export type AggregatedTime = {
    [key: string]: number; // key is the type, value is the total time in minutes
};

