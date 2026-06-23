import { Injectable, Inject, PLATFORM_ID, signal } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';

export type Theme = 'dark' | 'light';

@Injectable({
  providedIn: 'root',
})
export class ThemeService {
  private readonly THEME_KEY = 'fyp-scrum-theme';
  
  // Using an Angular Signal for easy, lightweight reactive template updates
  public currentTheme = signal<Theme>('dark');
  private isBrowser: boolean;

  constructor(@Inject(PLATFORM_ID) platformId: Object) {
    this.isBrowser = isPlatformBrowser(platformId);
    this.initializeTheme();
  }

  private initializeTheme(): void {
    if (this.isBrowser) {
      const savedTheme = localStorage.getItem(this.THEME_KEY) as Theme;
      const initialTheme = savedTheme === 'light' || savedTheme === 'dark' ? savedTheme : 'dark';
      this.setTheme(initialTheme);
    }
  }

  public setTheme(theme: Theme): void {
    if (this.isBrowser) {
      document.documentElement.setAttribute('data-theme', theme);
      localStorage.setItem(this.THEME_KEY, theme);
      this.currentTheme.set(theme);
    }
  }

  public toggleTheme(): void {
    const nextTheme = this.currentTheme() === 'dark' ? 'light' : 'dark';
    this.setTheme(nextTheme);
  }
}
