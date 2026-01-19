import { Inject, Injectable } from "@angular/core";
import { Position, CreatePositionDto } from "../../domain/models/position.model";
import { POSITION_REPOSITORY, PositionRepository } from "../../domain/ports/position.repo";

@Injectable({ providedIn: 'root' })
export class CreatePositionUseCase {
    constructor(@Inject(POSITION_REPOSITORY) private positionRepo: PositionRepository) {}
    
    async Execute(position: CreatePositionDto): Promise<Position> {
        return this.positionRepo.create(position);
    }
}
