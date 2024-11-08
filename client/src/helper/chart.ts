
function timeToMinutes(timeStr: string) {
    const [hours, minutes] = timeStr.split(":").map(Number);
    return hours * 60 + minutes;
}

export function aggregatedData(filteredTasks) {
    return filteredTasks.reduce(
        (acc, item) => {
            const type = item.type;
            const timeInMinutes = timeToMinutes(item.time);

            if (!acc[type]) {
                acc[type] = 0;
            }
            acc[type] += timeInMinutes;

            return acc
        }
    )
}

export function chartConfig(chartType, labels, data) {
    return ({
        type: chartType,
        data: {
            labels: labels,
            datasets: [
                {
                    label: "Time by Type (in minutes)",
                    data: data,
                    backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0', '#9966FF']
                }
            ]
        },
        options: {
            responsive: true,
            plugins: {
                legend: {
                    display: true
                },
                tooltip: {
                    callbacks: {
                        label: (tooltipItem) => {
                            const label = tooltipItem.label || '';
                            const value = tooltipItem.raw || 0;
                            return `${label}: ${value} minutes`;
                        }
                    }
                }
            }
        }
    })
}