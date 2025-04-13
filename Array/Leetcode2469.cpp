#include <vector>
std::vector<double> convertTemperature(double celsius) {
  std::vector<double> ans;
  double kelvin = celsius + 273.15;
  double Fahrenheit = celsius * 1.80 + 32.00;
  ans.push_back(kelvin);
  ans.push_back(Fahrenheit);
  return ans;
}
