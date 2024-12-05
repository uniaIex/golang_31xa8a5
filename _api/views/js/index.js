







// Basic GET request
async function fetchData(url, id_spinner) {
    try {
      displayLoading(id_spinner);
      const response = await fetch(url);
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      hideLoading(id_spinner); // Use .json() for JSON responses
      return data;
    } catch (error) {
      console.error('There was a problem with the fetch operation:', error);
    }
  }
  
  // Basic POST request with JSON data
async function postData(url, data) {
    try {
      displayLoading();
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
      hideLoading();
      return responseData;
    } catch (error) {
      console.error('There was a problem with the fetch operation:', error);
    }
  }

function buttons_reset()
{
    try {
        document.getElementById("button-responsive-projects").classList.remove("btn-dark");
        document.getElementById("button-responsive-technologies").classList.remove("btn-dark");
        console.log("removed btn classes")
      }
      catch(err) {
        console.log("could not reset interface");
      }
}



// showing loading
function displayLoading(id) {
    const loader = document.querySelector(id);
    loader.classList.remove("visually-hidden");
    // to stop loading after some time
    setTimeout(() => {
        loader.classList.remove("display");
    }, 5000);
}

// hiding loading 
function hideLoading(id) {
    const loader = document.querySelector(id);
    loader.classList.add("visually-hidden");
}

function getProjects()
{
    

    try {
        const elements = document.querySelectorAll('#reactive-elements');
        elements.forEach(element => element.remove());
      }
      catch(err) {
        console.log("`");
      }

    
    buttons_reset();
    document.getElementById("button-responsive-projects").classList.remove("btn-outline-dark");
    document.getElementById("button-responsive-projects").classList.add("btn-dark");
    
    

    fetchData('/api/project', '#project-loading-data')
    .then(data => {

        try {
          document.getElementById('responsive-projects-nr').textContent = 
        `${JSON.stringify(data.data.length)}`;
        } catch (error) {
          console.log('err');
        }

        data.data.forEach(project => {
            if (project.deployment != "-")
            {
              markup = 
            `
            <div class="card col-xs-3 border-dark m-3 bg-black text-white rounded-4 p-0" style="max-width: 25rem;" id="reactive-elements">
              
              <div class="card-header">${project.name}</div>
                
                <div class="card-body text-light p-3">
                  description:
                  <p class="card-text">${project.description}</p>
                  </div>
                  <div class="card-body text-light p-3">
                  technology:
                  <p class="card-text">${project.technology}</p>
                  </div>
                  <div class="card-body text-light p-3">
                  platform:
                  <p class="card-text">${project.platform}</p>
                  </div>
                  <div class="card-body text-light p-3">
                  
                  <a  class="text-light" href="${project.deployment}">${project.deployment}</a>
                  </div>
                  <div class="card-body text-light p-3">
                  <a href="${project.github}"><button type="button" class="btn btn-primary">GitHub</button></a>
                  </div>
            </div>
            `;
            }else
            {
             markup = 
            `
            <div class="card col-xs-3 border-dark m-3 bg-black text-white rounded-4 p-0" style="max-width: 25rem;" id="reactive-elements">
              
              <div class="card-header">${project.name}</div>
                
                <div class="card-body text-light p-3">
                  description:
                  <p class="card-text">${project.description}</p>
                  </div>
                  <div class="card-body text-light p-3">
                  technology:
                  <p class="card-text">${project.technology}</p>
                  </div>
                  <div class="card-body text-light p-3">
                  platform:
                  <p class="card-text">${project.platform}</p>
                  </div>
                  <div class="card-body text-light p-3">
                  <a href="${project.github}"><button type="button" class="btn btn-primary">GitHub</button></a>
                  </div>
            </div>
            `;
            }
              
            document.querySelector('#projects-view').insertAdjacentHTML('beforeend', markup)
        });

    })

    
    
}


function getTechnologies()
{
    

    try {
        const elements = document.querySelectorAll('#reactive-elements');
        elements.forEach(element => element.remove());
      }
      catch(err) {
        console.log("`");
      }

    buttons_reset();
    document.getElementById("button-responsive-technologies").classList.remove("btn-outline-dark");
    document.getElementById("button-responsive-technologies").classList.add("btn-dark");
    
    

    fetchData('/api/technology', '#project-loading-data')
    .then(data => {
        console.log(data);
        
        data.data.forEach(technology => {
          const markup = 
            `
            <div class="card col-xs-6 border-dark m-3 bg-black text-white p-0 rounded-4" style="max-width: 25rem;" id="reactive-elements">
              
              <div class="card-header justify-content-center">${technology.name}</div>
                <div class="card-body">
                  
                  <p class="card-text ">${technology.description}</p>
                  </div>
            </div>
            
            
            `
            document.querySelector('#projects-view').insertAdjacentHTML('beforeend', markup)
        });

    })
    
  
}

function getGraph()
{
    
    

    fetchData('/api/language', '#graph-loading-data')
    .then(data => {
        console.log(data);
        
        data.data.forEach(language => {
            const markup = 
            `
            
            <tr>
              <th scope="row"> ${language.name} </th>
              <td style="--size: ${language.value}; --color: rgb(202,167,196)"><span class="tooltip"> ${language.tooltip} </span></td>
            </tr>
            
            
            `
            document.querySelector('#reactive-graph').insertAdjacentHTML('beforeend', markup)
        });

    })
    //addFooter();
    
}