import { Component, EventEmitter, input, Output, signal, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Department } from '../../../../../core/domain/models';
import { CreateDepartmentUseCase, UpdateDepartmentUseCase } from '../../../../../core/usecases/department';

@Component({
  selector: 'app-department-form',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './department-form.component.html',
  styleUrl: './department-form.component.css'
})
export class DepartmentFormComponent implements OnInit {
  department = input<Department | null>(null);
  @Output() cancel = new EventEmitter<void>();
  @Output() saved = new EventEmitter<void>();

  private createDepartment = inject(CreateDepartmentUseCase);
  private updateDepartment = inject(UpdateDepartmentUseCase);

  departmentName = signal<string>('');

  ngOnInit(): void {
    const dept = this.department();
    if (dept) {
      this.departmentName.set(dept.name);
    }
  }

  updateName(event: Event) {
    const input = event.target as HTMLInputElement;
    this.departmentName.set(input.value);
  }

  isValid(): boolean {
    return this.departmentName().trim().length > 0;
  }

  close() {
    this.cancel.emit();
  }

  save() {
    if (!this.isValid()) return;

    const dept: Department = {
      id: this.department()?.id || '',
      name: this.departmentName()
    };

    const promise = this.department() 
      ? this.updateDepartment.Execute(dept)
      : this.createDepartment.Execute(dept);

    promise.then(() => {
      this.saved.emit();
    }).catch(err => {
      console.error('Error saving department:', err);
    });
  }
}
