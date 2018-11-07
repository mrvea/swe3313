import { Component } from '@angular/core';

const MENU_LIST = [
	{
		Name: "item",
		Path: ""
	},
	{
		Name: "item2",
		Path: ""
	},
	{
		Name: "item3",
		Path: ""
	}
];
@Component({
  selector: 'app-home',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
	MenuList: any[] = [];
 	Title = "Geb's Pizza";
 	constructor(){
 		this.MenuList = MENU_LIST;
 	}

 	UpdateWidth(){

 	}
	GetAnimation(o){

	}
}
