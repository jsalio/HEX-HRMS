import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { firstValueFrom } from "rxjs";
import { LoginUser, User, UserData } from "../core/domain/models";
import { UserRepository } from "../core/domain/ports/user.repo";

@Injectable({ providedIn: 'root' })
export class UserApiRepository implements UserRepository {

  constructor(private http: HttpClient) {}

 async login(loginUser: LoginUser): Promise<UserData> {
      const dto = await firstValueFrom( 
          this.http.post<UserData>(`http://localhost:5000/api/auth/login`, loginUser)
      );
      return dto;
  }

}
