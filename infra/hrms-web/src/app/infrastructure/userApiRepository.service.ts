import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { firstValueFrom } from "rxjs";
import { CreateUser, Filter, LoginUser, ModifyUser, PaginatedResponse, SearchQuery, User, UserData } from "../core/domain/models";
import { UserRepository } from "../core/domain/ports/user.repo";
import { environment } from "../../environments/environment.development";

@Injectable({ providedIn: 'root' })
export class UserApiRepository implements UserRepository {

  constructor(private http: HttpClient) {}

 async login(loginUser: LoginUser): Promise<UserData> {
      const dto = await firstValueFrom( 
          this.http.post<UserData>(`${environment.apiUrl}/auth/login`, loginUser)
      );
      return dto;
  }

  async list(query: SearchQuery): Promise<PaginatedResponse<UserData>> {
    const dto = await firstValueFrom(
      this.http.post<PaginatedResponse<UserData>>(`${environment.apiUrl}/auth/list`, query)
    );
    return dto;
  }

  async createUser(user: CreateUser): Promise<UserData> {
    const dto = await firstValueFrom(
      this.http.post<UserData>(`${environment.apiUrl}/auth/create`, user)
    );
    return dto;
  }

  async getUserByFilter(filter: Filter): Promise<UserData> {
    const dto = await firstValueFrom(
      this.http.post<UserData>(`${environment.apiUrl}/auth/get-user-by-field`, filter)
    );
    return dto;
  }

  async updateUser(user: ModifyUser): Promise<any> {
    const dto = await firstValueFrom(
      this.http.post<any>(`${environment.apiUrl}/auth/update`, user)
    );
    return dto;
  }

  

}
