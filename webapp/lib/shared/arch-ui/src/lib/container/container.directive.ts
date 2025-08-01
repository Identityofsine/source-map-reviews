import { Directive, inject, OnInit, Renderer2, ElementRef } from '@angular/core';

@Directive({
  selector: '[arch-container]',
  standalone: true,
})
export class ArchContainer implements OnInit {

  private readonly renderer = inject(Renderer2);
  private readonly elementRef = inject(ElementRef);

  constructor() { }

  public ngOnInit(): void {
    this.renderer.addClass(this.elementRef.nativeElement, 'arch-container');
  }

}
