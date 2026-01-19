import { Component, EventEmitter, input, Output, signal, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Position, CreatePositionDto, UpdatePositionDto, WorkType, PositionStatus } from '../../../../../core/domain/models/position.model';
import { Department } from '../../../../../core/domain/models';
import { CreatePositionUseCase, UpdatePositionUseCase } from '../../../../../core/usecases/position';

/**
 * Presenter component for position form (create/edit)
 */
@Component({
  selector: 'app-position-form',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './position-form.component.html',
  styleUrl: './position-form.component.css'
})
export class PositionFormComponent implements OnInit {
  position = input<Position | null>(null);
  departments = input<Department[]>([]);
  
  @Output() cancel = new EventEmitter<void>();
  @Output() saved = new EventEmitter<void>();

  private createPosition = inject(CreatePositionUseCase);
  private updatePosition = inject(UpdatePositionUseCase);

  // Form fields
  title = signal<string>('');
  code = signal<string>('');
  description = signal<string>('');
  requiredSkills = signal<string>('');
  salaryMin = signal<number>(0);
  salaryMax = signal<number>(0);
  maxEmployees = signal<number>(1);
  currency = signal<string>('USD');
  workType = signal<WorkType>('OnSite');
  departmentId = signal<string>('');
  status = signal<PositionStatus>('Active');
  
  saving = signal<boolean>(false);

  workTypes: { value: WorkType; label: string }[] = [
    { value: 'Remote', label: 'Remoto' },
    { value: 'Hybrid', label: 'Híbrido' },
    { value: 'OnSite', label: 'En Sitio' }
  ];

  statuses: { value: PositionStatus; label: string }[] = [
    { value: 'Active', label: 'Activo' },
    { value: 'Inactive', label: 'Inactivo' },
    { value: 'Closed', label: 'Cerrado' }
  ];

  ngOnInit(): void {
    const pos = this.position();
    if (pos) {
      this.title.set(pos.title);
      this.code.set(pos.code);
      this.description.set(pos.description);
      this.requiredSkills.set(pos.requiredSkills);
      this.salaryMin.set(pos.salaryMin);
      this.salaryMax.set(pos.salaryMax);
      this.maxEmployees.set(pos.maxEmployees);
      this.currency.set(pos.currency);
      this.workType.set(pos.workType);
      this.departmentId.set(pos.departmentId);
      this.status.set(pos.status);
    }
  }

  isValid(): boolean {
    return (
      this.title().trim().length > 0 &&
      this.code().trim().length > 0 &&
      this.departmentId().trim().length > 0 &&
      this.maxEmployees() > 0
    );
  }

  close() {
    this.cancel.emit();
  }

  save() {
    if (!this.isValid() || this.saving()) return;

    this.saving.set(true);

    if (this.position()) {
      const dto: UpdatePositionDto = {
        id: this.position()!.id,
        title: this.title(),
        code: this.code(),
        description: this.description(),
        requiredSkills: this.requiredSkills(),
        salaryMin: this.salaryMin(),
        salaryMax: this.salaryMax(),
        maxEmployees: this.maxEmployees(),
        currency: this.currency(),
        workType: this.workType(),
        departmentId: this.departmentId(),
        status: this.status()
      };

      this.updatePosition.Execute(dto)
        .then(() => this.saved.emit())
        .catch(err => console.error('Error updating position:', err))
        .finally(() => this.saving.set(false));
    } else {
      const dto: CreatePositionDto = {
        title: this.title(),
        code: this.code(),
        description: this.description(),
        requiredSkills: this.requiredSkills(),
        salaryMin: this.salaryMin(),
        salaryMax: this.salaryMax(),
        maxEmployees: this.maxEmployees(),
        currency: this.currency(),
        workType: this.workType(),
        departmentId: this.departmentId()
      };

      this.createPosition.Execute(dto)
        .then(() => this.saved.emit())
        .catch(err => console.error('Error creating position:', err))
        .finally(() => this.saving.set(false));
    }
  }
}
