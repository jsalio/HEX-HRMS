import { Inject, Injectable } from "@angular/core";
import { Permission } from "../../domain/models/permission.model";
import { ROLE_REPOSITORY, RoleRepository } from "../../domain/ports/role.repo";

@Injectable({ providedIn: 'root' })
export class ListSystemPermissionsUseCase {
    constructor(@Inject(ROLE_REPOSITORY) private roleRepo: RoleRepository) {}
    
    async Execute(): Promise<Permission[]> {
        return this.roleRepo.listSystemPermissions();
    }
}
