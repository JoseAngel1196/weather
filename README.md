# Weather CLI

A CLI tool that retrieves and displays the current weather conditions for a given location.

## Installation

To install the tool, simply run the following command:

```bash
go install github.com/yourusername/weather-cli
```

## Usage

To use the tool, run the following command:

```bash
weather-cli --location "New York"
```

Replace "New York" with the desired location.

By default, the tool will display the temperature in Fahrenheit. You can specify the temperature unit using the --unit flag, like so:

```bash
weather-cli --location "New York" --unit celsius
```

The available temperature units are kelvin, celsius, and fahrenheit.

## Contributing

Contributions are always welcome! If you'd like to contribute to the project, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
