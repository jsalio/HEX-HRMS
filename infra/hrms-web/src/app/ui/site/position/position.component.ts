import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListPositionUseCase, DeletePositionUseCase } from '../../../core/usecases/position';
import { ListDepartmentUseCase } from '../../../core/usecases/department';
import { Position } from '../../../core/domain/models/position.model';
import { Department } from '../../../core/domain/models';
import { PositionListComponent } from './components/position-list/position-list.component';
import { PositionFormComponent } from './components/position-form/position-form.component';

/**
 * Container component for Position management
 */
@Component({
  selector: 'app-position',
  standalone: true,
  imports: [CommonModule, PositionListComponent, PositionFormComponent],
  templateUrl: './position.component.html',
  styleUrl: './position.component.css'
})
export class PositionComponent implements OnInit {
  private listPosition = inject(ListPositionUseCase);
  private deletePositionUseCase = inject(DeletePositionUseCase);
  private listDepartmentUseCase = inject(ListDepartmentUseCase);

  positions = signal<Position[]>([]);
  departments = signal<Department[]>([]);
  
  // Modal State
  showForm = signal<boolean>(false);
  selectedPosition = signal<Position | null>(null);
  loading = signal<boolean>(false);

  ngOnInit(): void {
    this.fetchPositions();
    this.fetchDepartments();
  }

  fetchPositions(): void {
    this.loading.set(true);
    this.listPosition.Execute({
      filters: [],
      pagination: { page: 1, limit: 100 }
    }).then((data) => {
      this.positions.set(data.rows || []);
    }).catch(err => {
      console.error('Error fetching positions:', err);
    }).finally(() => {
      this.loading.set(false);
    });
  }

  fetchDepartments(): void {
    this.listDepartmentUseCase.Execute({
      filters: [],
      pagination: { page: 1, limit: 100 }
    }).then((data) => {
      this.departments.set(data.rows || []);
    }).catch(err => {
      console.error('Error fetching departments:', err);
    });
  }

  // --- Actions ---

  createPosition() {
    this.selectedPosition.set(null);
    this.showForm.set(true);
  }

  editPosition(position: Position) {
    this.selectedPosition.set(position);
    this.showForm.set(true);
  }

  deletePosition(id: string) {
    if (!confirm('¿Está seguro de eliminar esta posición?')) return;
    
    this.deletePositionUseCase.Execute(id).then(() => {
      this.fetchPositions();
    }).catch(err => {
      console.error('Error deleting position:', err);
    });
  }

  closeForm() {
    this.showForm.set(false);
    this.selectedPosition.set(null);
  }

  onSaved() {
    this.fetchPositions();
    this.closeForm();
  }
}
