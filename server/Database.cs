using Microsoft.EntityFrameworkCore;

class TasksDb : DbContext
{
    public TasksDb(DbContextOptions<TasksDb> options)
        : base(options) { }

    public DbSet<Task> Tasks => Set<Task>();
}