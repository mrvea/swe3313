import { Component, OnInit } from '@angular/core';
import { Pizza, DoughTypes, ToppingTypes, SauceTypes, Product } from '../product/models';

const temp_pizzas = [
	new Pizza({
		ID: 1,
		Name: "Cheese Pizza",
		Price: 15.99,
		Dough: {
			ID: 11,
			Name: "Thin Crush",
			Type:  DoughTypes.THIN,
			Price: 4,
		},
		Sauce: {
			ID: 12,
			Name: "Marinara",
			Type: SauceTypes.RED,
			Price: 2
		},
		Toppings: [
			{
				ID: 13,
				Name: "House Blend Cheese",
				Tooltip: "mozzarella, provolone, cheddar and Parmesan",
				Type: ToppingTypes.CHEESE
			}
		]
	}),
	new Pizza({
		ID: 2,
		Name: "Pepperoni Pizza",
		Price: 17.99,
		Image: "pizza_pepperoni.png",
		Dough: {
			ID: 11,
			Name: "Pan Crust",
			Type:  DoughTypes.PAN,
			Price: 4,
		},
		Sauce: {
			ID: 12,
			Name: "Marinara",
			Type: SauceTypes.RED,
			Price: 2
		},
		Toppings: [
			{
				ID: 23,
				Name: "House Blend Cheese",
				Tooltip: "mozzarella, provolone, cheddar and Parmesan",
				Type: ToppingTypes.CHEESE
			},
			{
				ID: 24,
				Name: "Pepperoni",
				Type: ToppingTypes.MEAT
			}
		]
	}),
	new Pizza({
		ID: 3,
		Name: "Meat Lovers Pizza",
		Price: 19.99,
		Image: "pan_meat_lovers.png",
		Dough: {
			ID: 11,
			Name: "Pan Crust",
			Type:  DoughTypes.PAN,
			Price: 4,
		},
		Sauce: {
			ID: 12,
			Name: "Marinara",
			Type: SauceTypes.RED,
			Price: 2
		},
		Toppings: [
			{
				ID: 13,
				Name: "House Blend Cheese",
				Tooltip: "mozzarella, provolone, cheddar and Parmesan",
				Type: ToppingTypes.CHEESE
			},
			{
				ID: 24,
				Name: "Pepperoni",
				Type: ToppingTypes.MEAT
			},
			{
				ID: 35,
				Name: "Sausage",
				Type: ToppingTypes.MEAT
			},
			{
				ID: 36,
				Name: "Ham",
				Type: ToppingTypes.MEAT
			},
			{
				ID: 37,
				Name: "Bacon",
				Type: ToppingTypes.MEAT
			}
		]
	})
]

@Component({
  selector: 'app-selection',
  templateUrl: './selection.component.html',
  styleUrls: ['./selection.component.css']
})
export class SelectionComponent implements OnInit {
	pizzas: Pizza[];
	products: Product[];
	stage: any;

	stages = [
		{
			stage: "menu",
			title: "Please suggest pre-made pizza, or make a custom.",
			fn: "getMenu"
		},
		{	
			stage: "crust",
			title: "Please ask about any updates to the crust and/or sause.",
			fn: "getCrustSizeSause"
		},
		{	
			stage: "toppings",
			title: "Please ask about any additional toppings.",
			fn: "getToppings"
		}
	]
  constructor() { 
  	this.pizzas = temp_pizzas;
  	this.setStage(0);
  }

  ngOnInit() {
  }

  setStage(stageIndex){
  	this.stage = this.stages[stageIndex];
  	this.stage.fn();
  }

  getMenu(){
  	this.products = temp_pizzas;
  }

  getCrustSizeSause(){
  	this.products = [];
  }

  getToppings(){
  	this.products = [];
  }
}
