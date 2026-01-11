import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { UserData } from '../../../../../core/domain/models';

@Component({
  selector: 'app-user-list',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './user-list.component.html',
  styleUrl: './user-list.component.css'
})
export class UserListComponent {
  @Input() users: UserData[] = [];
  @Input() currentPage: number = 1;
  @Input() totalPages: number = 1;

  @Output() pageChange = new EventEmitter<number>();
  @Output() delete = new EventEmitter<string>();
  @Output() toggleStatus = new EventEmitter<string>();

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
}
