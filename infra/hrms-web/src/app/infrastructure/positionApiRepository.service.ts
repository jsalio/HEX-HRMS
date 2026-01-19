import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { firstValueFrom } from "rxjs";
import { PaginatedResponse, SearchQuery } from "../core/domain/models";
import { Position, CreatePositionDto, UpdatePositionDto } from "../core/domain/models/position.model";
import { PositionRepository } from "../core/domain/ports/position.repo";
import { environment } from "../../environments/environment.development";

/**
 * HTTP implementation of PositionRepository
 */
@Injectable({ providedIn: 'root' })
export class PositionApiRepository implements PositionRepository {

  constructor(private http: HttpClient) {}

  /**
   * List positions with pagination and filters
   */
  async list(query: SearchQuery): Promise<PaginatedResponse<Position>> {
    const dto = await firstValueFrom(
      this.http.post<PaginatedResponse<Position>>(`${environment.apiUrl}/position/get-all`, query)
    );
    return dto;
  }

  /**
   * Get a position by ID
   */
  async get(id: string): Promise<Position> {
    const dto = await firstValueFrom(
      this.http.post<Position>(`${environment.apiUrl}/position/get`, { id })
    );
    return dto;
  }

  /**
   * Create a new position
   */
  async create(position: CreatePositionDto): Promise<Position> {
    const dto = await firstValueFrom(
      this.http.post<Position>(`${environment.apiUrl}/position/create`, position)
    );
    return dto;
  }

  /**
   * Update an existing position
   */
  async update(position: UpdatePositionDto): Promise<Position> {
    const dto = await firstValueFrom(
      this.http.post<Position>(`${environment.apiUrl}/position/update`, position)
    );
    return dto;
  }

  /**
   * Delete a position by ID
   */
  async delete(id: string): Promise<void> {
    await firstValueFrom(
      this.http.post(`${environment.apiUrl}/position/delete`, { id })
    );
  }
}
