import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvitedUserComponent } from './invited-user.component';

describe('InvitedUserComponent', () => {
  let component: InvitedUserComponent;
  let fixture: ComponentFixture<InvitedUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [InvitedUserComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(InvitedUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
