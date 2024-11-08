using Microsoft.EntityFrameworkCore;
using System;
using Npgsql;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddDbContext<TasksDb>(opt => opt.UseInMemoryDatabase("Timetracker Database"));
builder.Services.AddDatabaseDeveloperPageExceptionFilter();
var app = builder.Build();

app.MapTaskRoutes();


app.Run();
