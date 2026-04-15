import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { TaskService, Task } from '../../services/task';

@Component({
  selector: 'app-tasks',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css'
})
export class TasksComponent implements OnInit {
  tasks: Task[] = [];

  newTask: Task = { title: '', description: '' };

  
  editingTask: Task | null = null;

  constructor(private taskService: TaskService) {}


  ngOnInit() {
    this.loadTasks();
  }

  loadTasks() {
    this.taskService.getTasks().subscribe(tasks => {
      this.tasks = tasks;
    });
  }

  createTask() {
    if (!this.newTask.title) return;
    this.taskService.createTask(this.newTask).subscribe(() => {
      this.newTask = { title: '', description: '' };
      this.loadTasks();
    });
  }

  startEdit(task: Task) {
    this.editingTask = { ...task };
  }

  
  saveEdit() {
    if (!this.editingTask?.id) return;
    this.taskService.updateTask(this.editingTask.id, this.editingTask).subscribe(() => {
      this.editingTask = null;
      this.loadTasks();
    });
  }

  
  cancelEdit() {
    this.editingTask = null;
  }

  
  deleteTask(id: number) {
    this.taskService.deleteTask(id).subscribe(() => {
      this.loadTasks();
    });
  }

  
  toggleDone(task: Task) {
    this.taskService.updateTask(task.id!, { ...task, done: !task.done }).subscribe(() => {
      this.loadTasks();
    });
  }
}