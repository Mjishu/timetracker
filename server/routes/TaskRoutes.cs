using Microsoft.AspNetCore.Builder;
using Microsoft.EntityFrameworkCore;
using Microsoft.AspNetCore.Http;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;

public static class TaskRoutes
{
    public static void MapTaskRoutes(this IEndpointRouteBuilder routes)
    {
        var taskRoutes = routes.MapGroup("/tasks");

        //get
        taskRoutes.MapGet("/all", async (TasksDb db) =>
            await db.Tasks.ToListAsync());

        taskRoutes.MapGet("/{id}", async (int id, TasksDb db) =>
            await db.Tasks.FindAsync(id)
                is Task task
                    ? Results.Ok(task)
                    : Results.NotFound());


        //post
        taskRoutes.MapPost("/new", async (Task task, TasksDb db) =>
            {
                Console.WriteLine(task);
                db.Tasks.Add(task);
                await db.SaveChangesAsync();

                return Results.Created($"/tasks/{task.Id}", task);
            });

    }
}