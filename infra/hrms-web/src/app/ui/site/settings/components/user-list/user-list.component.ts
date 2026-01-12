import { Component, EventEmitter, Input, OnInit, output, Output, inject, signal, computed } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { UserData, Role } from '../../../../../core/domain/models';
import { ListRoleUseCase } from '../../../../../core/usecases/role';

@Component({
  selector: 'app-user-list',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './user-list.component.html',
  styleUrl: './user-list.component.css'
})
export class UserListComponent implements OnInit {
  @Input() users: UserData[] = [];
  @Input() currentPage: number = 1;
  @Input() totalPages: number = 1;

  @Output() pageChange = new EventEmitter<number>();
  @Output() delete = new EventEmitter<string>();
  @Output() toggleStatus = new EventEmitter<string>();
  edit = output<string>();

  private listRoleUseCase = inject(ListRoleUseCase);

  // Map of roleId -> roleName for quick lookup
  roleMap = signal<Map<string, string>>(new Map());

  ngOnInit(): void {
    this.fetchRoles();
  }

  fetchRoles(): void {
    this.listRoleUseCase.Execute({
      filters: [],
      pagination: { page: 1, limit: 100 }
    }).then(data => {
      const map = new Map<string, string>();
      data.rows.forEach(role => {
        map.set(role.id, role.name);
      });
      this.roleMap.set(map);
    }).catch(err => {
      console.error('Error fetching roles:', err);
    });
  }

  getRoleName(roleId: string): string {
    return this.roleMap().get(roleId) || roleId;
  }

  getInitials(user: UserData): string {
    if (!user.name || !user.lastName) return '??';
    return (user.name[0] + user.lastName[0]).toUpperCase();
  }

  onPageChange(page: number): void {
    if (page >= 1 && page <= this.totalPages) {
      this.pageChange.emit(page);
    }
  }

  onDelete(id: string): void {
    this.delete.emit(id);
  }

  onToggleStatus(id: string): void {
    this.toggleStatus.emit(id);
  }

  onEdit(id: string): void {
    this.edit.emit(id);
  }
}
