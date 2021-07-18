import {Component, OnInit} from '@angular/core';
import {PageService} from "./services/page.service";
import {PageModel} from "./models/page.model";

import {isUri, isHttpsUri, isHttpUri, isWebUri} from 'valid-url';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  pages: PageModel[] = [];

  blocked = true;

  lastAdded!: PageModel;

  newPageUrl!: string;

  constructor(private pageService: PageService) {
  }


  ngOnInit() {
    this.fetchPages();
  }

  onModelChange(): void {
    if (this.newPageUrl.startsWith("https://") || this.newPageUrl.startsWith("http://")) {

      this.blocked = !isUri(this.newPageUrl);

      return;
    }

    this.blocked = true;
  }

  async fetchPages(): Promise<void> {
    this.pages = await this.pageService.index();
  }

  async createPage(): Promise<void> {
    this.lastAdded = await this.pageService.create(this.newPageUrl);

    this.fetchPages();

    this.newPageUrl = null!;
  }
}
