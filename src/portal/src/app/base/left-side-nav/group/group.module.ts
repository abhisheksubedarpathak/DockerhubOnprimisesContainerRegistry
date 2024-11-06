import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { SharedModule } from '../../../shared/shared.module';
import { GroupComponent } from './group.component';
import { AddGroupModalComponent } from './add-group-modal/add-group-modal.component';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
    {
        path: '',
        component: GroupComponent,
    },
];
@NgModule({
    imports: [SharedModule, RouterModule.forChild(routes), FormsModule],
    declarations: [GroupComponent, AddGroupModalComponent],
})
export class GroupModule {}
