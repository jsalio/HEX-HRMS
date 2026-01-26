import { Component, EventEmitter, input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Position } from '../../../../../core/domain/models/position.model';

/**
 * Presenter component for displaying position list
 */
@Component({
  selector: 'app-position-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './position-list.component.html',
  styleUrl: './position-list.component.css'
})
export class PositionListComponent {
  positions = input<Position[]>([]);
  loading = input<boolean>(false);
  
  @Output() edit = new EventEmitter<Position>();
  @Output() delete = new EventEmitter<string>();

  getWorkTypeLabel(workType: string): string {
    const labels: Record<string, string> = {
      'Remote': 'Remoto',
      'Hybrid': 'Híbrido',
      'OnSite': 'En Sitio'
    };
    return labels[workType] || workType;
  }

  getStatusLabel(status: string): string {
    const labels: Record<string, string> = {
      'Active': 'Activo',
      'Inactive': 'Inactivo',
      'Closed': 'Cerrado'
    };
    return labels[status] || status;
  }

  getStatusClass(status: string): string {
    const classes: Record<string, string> = {
      'Active': 'status-active',
      'Inactive': 'status-inactive',
      'Closed': 'status-closed'
    };
    return classes[status] || '';
  }
}
