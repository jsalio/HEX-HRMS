import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListDepartmentUseCase, DeleteDepartmentUseCase } from '../../../core/usecases/department';
import { Department } from '../../../core/domain/models';
import { DepartmentListComponent } from './components/department-list/department-list.component';
import { DepartmentFormComponent } from './components/department-form/department-form.component';

@Component({
  selector: 'app-department',
  standalone: true,
  imports: [CommonModule, DepartmentListComponent, DepartmentFormComponent],
  templateUrl: './department.component.html',
  styleUrl: './department.component.css'
})
export class DepartmentComponent implements OnInit {
  private listDepartment = inject(ListDepartmentUseCase);
  private deleteDepartmentUseCase = inject(DeleteDepartmentUseCase);

  departments = signal<Department[]>([]);
  
  // Modal State
  showForm = signal<boolean>(false);
  selectedDepartment = signal<Department | null>(null);

  ngOnInit(): void {
    this.fetchDepartments();
  }

  fetchDepartments(): void {
    this.listDepartment.Execute({
      filters: [],
      pagination: { page: 1, limit: 100 }
    }).then((data) => {
      this.departments.set(data.rows);
    }).catch(err => {
      console.error('Error fetching departments:', err);
    });
  }

  // --- Actions ---

  createDepartment() {
    this.selectedDepartment.set(null);
    this.showForm.set(true);
  }

  editDepartment(department: Department) {
    this.selectedDepartment.set(department);
    this.showForm.set(true);
  }

  deleteDepartment(id: string) {
    this.deleteDepartmentUseCase.Execute(id).then(() => {
      this.fetchDepartments();
    }).catch(err => {
      console.error('Error deleting department:', err);
    });
  }

  closeForm() {
    this.showForm.set(false);
    this.selectedDepartment.set(null);
  }

  onSaved() {
    this.fetchDepartments();
    this.closeForm();
  }
}
