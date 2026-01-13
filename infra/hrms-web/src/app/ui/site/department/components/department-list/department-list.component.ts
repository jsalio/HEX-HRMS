import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Department } from '../../../../../core/domain/models';

@Component({
  selector: 'app-department-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './department-list.component.html',
  styleUrl: './department-list.component.css'
})
export class DepartmentListComponent {
  @Input() departments: Department[] = [];
  @Output() create = new EventEmitter<void>();
  @Output() edit = new EventEmitter<Department>();
  @Output() delete = new EventEmitter<string>();

  createNew() {
    this.create.emit();
  }

  editDepartment(department: Department) {
    this.edit.emit(department);
  }

  deleteDepartment(id: string) {
    this.delete.emit(id);
  }
}
