# 📖 Booklooker

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![Linters](https://github.com/Plaenkler/Booklooker/actions/workflows/linters.yml/badge.svg)](https://github.com/Plaenkler/Booklooker/actions/workflows/linters.yml)
[![Support me](https://img.shields.io/badge/Support%20me%20%E2%98%95-orange.svg)](https://www.buymeacoffee.com/Plaenkler)

Plaenkler/Booklooker enables HTTP requests to be made to the Booklooker API. The library abstracts the constructs and allows developers to interact with the Booklooker API through a convenient and intuitive interface.

<table style="border:none;">
  <tr>
    <td><img src="https://user-images.githubusercontent.com/60503970/217089943-53d1a509-6dd2-4fa0-85b4-380cc8a97530.png" width="480"/></td>
    <td><img src="https://user-images.githubusercontent.com/60503970/217089956-f6949215-ca0c-44f5-a3f8-73b736d8832d.png" width="480"/></td>
  </tr>
</table>

## 🎯 Project goals

- [x] Implementation of all models and handlers
- [x] Client with automatic token renewal
- [x] Implementation of one example for each handler

## ✅ Examples

A separate example has been implemented for each interface of the API. All examples use the client for authentication except the authentication example itself if there is a desire to implement a separate client.

- [Immediate deletion of a single item](https://github.com/Plaenkler/Booklooker/tree/main/examples/article)
- [Download all active part numbers](https://github.com/Plaenkler/Booklooker/tree/main/examples/article_list)
- [Query the status of an item](https://github.com/Plaenkler/Booklooker/tree/main/examples/article_status)
- [Authentication via API Key](https://github.com/Plaenkler/Booklooker/tree/main/examples/authenticate)
- [Upload quote or image files](https://github.com/Plaenkler/Booklooker/tree/main/examples/file_import)
- [Query the status of an uploaded quote file](https://github.com/Plaenkler/Booklooker/tree/main/examples/file_status)
- [Download of all orders of a time range](https://github.com/Plaenkler/Booklooker/tree/main/examples/order)
- [Canceling a complete order](https://github.com/Plaenkler/Booklooker/tree/main/examples/order_cancel)
- [Canceling the order of a single item](https://github.com/Plaenkler/Booklooker/tree/main/examples/order_item_cancel)
- [Sending a message to the customer](https://github.com/Plaenkler/Booklooker/tree/main/examples/order_message)
- [Setting the status of an order](https://github.com/Plaenkler/Booklooker/tree/main/examples/order_status)
