// Simple Node.js demo using axios
const axios = require("axios");

const API_KEY = process.env.API_KEY;
if (!API_KEY) {
  console.error("Please set the API_KEY environment variable.");
  process.exit(1);
}

// Replace with real SDK client initialization if available
const ENDPOINT = "https://api.example.com/v1/demo"; // placeholder

async function main() {
  try {
    const resp = await axios.post(
      ENDPOINT,
      { message: "hello from node demo" },
      { headers: { Authorization: `Bearer ${API_KEY}`, "Content-Type": "application/json" }, timeout: 10000 }
    );
    console.log("Response status:", resp.status);
    console.log("Body:", resp.data);
  } catch (err) {
    if (err.response) {
      console.error("Error response:", err.response.status, err.response.data);
    } else {
      console.error("Request error:", err.message);
    }
    process.exit(1);
  }
}

main();
