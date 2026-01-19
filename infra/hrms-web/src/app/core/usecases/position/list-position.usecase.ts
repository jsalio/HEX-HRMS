import { Inject, Injectable } from "@angular/core";
import { PaginatedResponse, SearchQuery } from "../../domain/models";
import { Position } from "../../domain/models/position.model";
import { POSITION_REPOSITORY, PositionRepository } from "../../domain/ports/position.repo";

@Injectable({ providedIn: 'root' })
export class ListPositionUseCase {
    constructor(@Inject(POSITION_REPOSITORY) private positionRepo: PositionRepository) {}
    
    async Execute(query: SearchQuery): Promise<PaginatedResponse<Position>> {
        return this.positionRepo.list(query);
    }
}
