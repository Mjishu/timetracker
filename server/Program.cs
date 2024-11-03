using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddDbContext<TasksDb>(opt => opt.UseInMemoryDatabase("Timetracker Database"));
builder.Services.AddDatabaseDeveloperPageExceptionFilter();
var app = builder.Build();

//todo Figure out how to get these into their own module and then import that module into program.cs!
app.MapGet("/tasks", async (TasksDb db) =>
    await db.Tasks.ToListAsync());

app.MapPost("/tasks", async (Task task, TasksDb db) =>
{
    db.Tasks.Add(task);
    await db.SaveChangesAsync();

    return Results.Created($"/tasks/{task.Id}", task);
});

app.Run();
