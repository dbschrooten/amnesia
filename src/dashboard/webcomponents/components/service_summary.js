'use strict';

import { LitElement, html } from 'lit-element';
import ApexCharts from 'apexcharts';

class ServiceSummary extends LitElement {

  constructor() {
    super();
    this.options = {
      series: [{
          name: "Desktops",
          data: [10, 41, 35, 51, 49, 62, 69, 91, 148]
      }],
      chart: {
        sparkline: {
          enabled: true
        },
        dropShadow: {
          enabled: true,
          // enabledOnSeries: ['series-1'],
          // top: 0,
          // left: 0,
          // blur: 3,
          // color: '#000000',
          opacity: 0.1
        }
      },
      dataLabels: {
        enabled: false
      },
      stroke: {
        curve: 'smooth',
        width: 2
      },
      colors: ['#c53030'],
      grid: {
        row: {
          colors: ['#f3f3f3', 'transparent'], // takes an array which will be repeated on columns
          opacity: 0.5
        },
      },
      xaxis: {
        categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep'],
      },
      plotOptions: {
        line: {
          columnWidth: '20%',
          startingShape: 'rounded',
          endingShape: 'rounded'
        }
      }
    };
  }

  renderChart() {
    console.log('Render chart');
    const el = document.createElement('div');
    const chart = new ApexCharts(el, this.options);
    this.shadowRoot.querySelector('#service-chart').appendChild(el);
    chart.render();
  }

  firstUpdated() {
    this.renderChart();
  }

  render() {
    return html`
      <div id="service-chart"></div>
      <slot></slot>
    `;
  }

}

customElements.define('service-summary', ServiceSummary);
