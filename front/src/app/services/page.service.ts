import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {PageModel} from "../models/page.model";

@Injectable({
  providedIn: 'root'
})
export class PageService {

  private readonly apiUrl = 'http://localhost:3000'

  constructor(private httpClient: HttpClient) {
  }

  async index(): Promise<PageModel[]> {
    return this.httpClient.get<PageModel[]>(this.apiUrl + '/pages').toPromise();
  }

  async create(url: string): Promise<PageModel> {
    return this.httpClient.post<PageModel>(this.apiUrl + '/pages', { url }).toPromise();
  }
}
