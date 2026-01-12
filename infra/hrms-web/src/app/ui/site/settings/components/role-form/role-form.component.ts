import { Component, computed, EventEmitter, inject, input, OnInit, Output, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Role, Permission } from '../../../../../core/domain/models';
import { ListSystemPermissionsUseCase, CreateRoleUseCase, UpdateRoleUseCase } from '../../../../../core/usecases/role';

@Component({
  selector: 'app-role-form',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './role-form.component.html',
  styleUrl: './role-form.component.css'
})
export class RoleFormComponent implements OnInit {
  role = input<Role | null>(null);
  @Output() cancel = new EventEmitter<void>();
  @Output() saved = new EventEmitter<void>();

  private listPermissions = inject(ListSystemPermissionsUseCase);
  private createRole = inject(CreateRoleUseCase);
  private updateRole = inject(UpdateRoleUseCase);

  roleName = signal<string>('');
  
  // All system permissions fetched from backend
  systemPermissions = signal<Permission[]>([]);

  // Permissions currently assigned to the role
  assignedPermissions = signal<Permission[]>([]);

  // Permissions available to be assigned (System - Assigned)
  availablePermissions = computed(() => {
    const assignedIds = new Set(this.assignedPermissions().map(p => p.name));
    return this.systemPermissions().filter(p => !assignedIds.has(p.name));
  });

  // Selection states for Transfer Panel
  selectedAvailable = signal<Permission[]>([]);
  selectedAssigned = signal<Permission[]>([]);

  // Search filters
  searchAvailableTerm = signal<string>('');
  searchAssignedTerm = signal<string>('');

  filteredAvailablePermissions = computed(() => {
    const term = this.searchAvailableTerm().toLowerCase();
    return this.availablePermissions().filter(p => 
      p.name.toLowerCase().includes(term) || p.description.toLowerCase().includes(term)
    );
  });

  filteredAssignedPermissions = computed(() => {
    const term = this.searchAssignedTerm().toLowerCase();
    return this.assignedPermissions().filter(p => 
      p.name.toLowerCase().includes(term) || p.description.toLowerCase().includes(term)
    );
  });

  ngOnInit(): void {
    // Load system permissions
    this.listPermissions.Execute().then(perms => {
      this.systemPermissions.set(perms);
      
      // If editing, initialize form
      const role = this.role();
      if (role) {
        this.roleName.set(role.name);
        // Ensure we match objects from system permissions to keep consistency if needed, 
        // or just use the role permissions if they match the interface.
        this.assignedPermissions.set(role.permissions);
      }
    });

  }

  updateName(event: Event) {
    const input = event.target as HTMLInputElement;
    this.roleName.set(input.value);
  }

  // --- Transfer Panel Logic ---

  selectAvailable(perm: Permission) {
    this.selectedAvailable.update(selected => {
      if (selected.includes(perm)) {
        return selected.filter(p => p !== perm);
      } else {
        return [...selected, perm];
      }
    });
  }

  isSelectedAvailable(perm: Permission): boolean {
    return this.selectedAvailable().includes(perm);
  }

  selectAssigned(perm: Permission) {
    this.selectedAssigned.update(selected => {
      if (selected.includes(perm)) {
        return selected.filter(p => p !== perm);
      } else {
        return [...selected, perm];
      }
    });
  }

  isSelectedAssigned(perm: Permission): boolean {
    return this.selectedAssigned().includes(perm);
  }

  moveToAssigned() {
    const toMove = this.selectedAvailable();
    this.assignedPermissions.update(assigned => [...assigned, ...toMove]);
    this.selectedAvailable.set([]); // Clear selection
  }

  moveToAvailable() {
    const toMove = this.selectedAssigned();
    // Removing from assigned implicitly moves to available because of the computed `availablePermissions`
    this.assignedPermissions.update(assigned => assigned.filter(p => !toMove.includes(p)));
    this.selectedAssigned.set([]); // Clear selection
  }

  moveAllToAssigned() {
    this.assignedPermissions.set([...this.systemPermissions()]);
    this.selectedAvailable.set([]);
  }

  moveAllToAvailable() {
    this.assignedPermissions.set([]);
    this.selectedAssigned.set([]);
  }

  searchAvailable(event: Event) {
    const input = event.target as HTMLInputElement;
    this.searchAvailableTerm.set(input.value);
  }

  searchAssigned(event: Event) {
    const input = event.target as HTMLInputElement;
    this.searchAssignedTerm.set(input.value);
  }

  // --- Save Logic ---

  isValid(): boolean {
    return this.roleName().trim().length > 0;
  }

  close() {
    this.cancel.emit();
  }

  save() {
    if (!this.isValid()) return;

    const roleData: Role = {
      id: this.role()?.id || '',
      name: this.roleName(),
      permissions: this.assignedPermissions()
    };

    const promise = this.role() 
      ? this.updateRole.Execute(roleData)
      : this.createRole.Execute(roleData);

    promise.then(() => {
      this.saved.emit();
    }).catch(err => {
      console.error('Error saving role:', err);
      // Handle error (toast notification, etc.)
    });
  }
}
