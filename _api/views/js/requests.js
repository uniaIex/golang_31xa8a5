// Basic GET request
async function fetchData(url) {
    try {
      const response = await fetch(url);
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json(); // Use .json() for JSON responses
      return data;
    } catch (error) {
      console.error('There was a problem with the fetch operation:', error);
    }
  }
  
  // Basic POST request with JSON data
async function postData(url, data) {
    try {
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const responseData = await response.json();
      return responseData;
    } catch (error) {
      console.error('There was a problem with the fetch operation:', error);
    }
  }
/*  
  // Example usage
  // GET request
  fetchData('https://your-server.com/api/data')
    .then(data => console.log(data));
  
  // POST request
  postData('https://your-server.com/api/submit', { key: 'value' })
    .then(data => console.log(data));
*/