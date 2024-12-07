provider "beverage" {
  base_url = "http://localhost:5000"
}

resource "beverage" "example" {
  name = "Coke"
  type = "Soda"
}

