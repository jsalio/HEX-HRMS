import { Component, EventEmitter, Output, input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Role } from '../../../../../core/domain/models';

@Component({
  selector: 'app-role-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './role-list.component.html',
  styleUrl: './role-list.component.css'
})
export class RoleListComponent {
  roles = input.required<Role[]>();
  @Output() create = new EventEmitter<void>();
  @Output() edit = new EventEmitter<Role>();

  createNewRole() {
    this.create.emit();
  }

  editRole(role: Role) {
    this.edit.emit(role);
  }
}
