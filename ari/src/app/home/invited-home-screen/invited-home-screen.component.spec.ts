import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvitedHomeScreenComponent } from './invited-home-screen.component';

describe('InvitedHomeScreenComponent', () => {
  let component: InvitedHomeScreenComponent;
  let fixture: ComponentFixture<InvitedHomeScreenComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [InvitedHomeScreenComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(InvitedHomeScreenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
