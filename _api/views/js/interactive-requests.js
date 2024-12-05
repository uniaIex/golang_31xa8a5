
   // Async function to handle form submission without page reload
   async function submitFormData(event) {
    // Prevent the default form submission which would reload the page
    event.preventDefault();
    
    // Get form data
    const form = event.target;
    const formData = new FormData(form);
    //var data = Object.fromEntries(formData.entries());
  
    const data = {
      // Use Number() to convert to a number
      // Alternatively, use parseFloat() or parseInt() depending on your needs
      name: formData.get('name'),
      description: formData.get('description'),
      github: formData.get('github'),
      deployment: formData.get('deployment'),
      technology: Number(formData.get('technology')),
      platform: formData.get('platform')
    };
    
    
    try {
      // Send data to server
      const response = await fetch('/api/project/add', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
      });
      
      // Check if request was successful
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      
      // Parse the response
      const responseData = await response.json();
      
      // Update page content dynamically without refreshing
      document.getElementById('result-container').textContent = 
        `Received: ${JSON.stringify(responseData)}`;
      
    } catch (error) {
      // Handle any errors
      console.error('Error:', error);
      document.getElementById('result-container').textContent = 
        `Error occurred: ${error.message}`;
    }
  }
  
  // Example HTML structure this would work with
  /*
  <form id="dynamic-form" onsubmit="submitFormData(event)">
    <input type="text" name="username">
    <input type="email" name="email">
    <button type="submit">Send</button>
  </form>
  <div id="result-container"></div>
  */
  
  // Real-time data fetching example
  async function fetchLiveData() {
    try {
      const response = await fetch('/api/project');
      const data = await response.json();
      
      // Update specific parts of the page
      document.getElementById('live-data-display').innerHTML = `
        Current value: ${JSON.stringify(data)}
      `;
    } catch (error) {
      console.error('Failed to fetch live data:', error);
    }
  }
  
  // Optionally set up periodic updates
  setInterval(fetchLiveData, 1000);