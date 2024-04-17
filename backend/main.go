package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	endpoints "example.com/backend/Endpoints"
)



func main(){
	router:=gin.Default()
	router.POST("/User", func(c *gin.Context) {
        var user endpoints.Users
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        id, err := endpoints.AddUser(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"id": id})
    })

    router.GET("/User", func(ctx *gin.Context) {
        users, err := endpoints.FetchUsers()
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusOK, users)
    })

    router.PUT("/User/:id", func(ctx *gin.Context) {
       
        userID := ctx.Param("id")
    
        var updatedUser endpoints.Users
        if err := ctx.BindJSON(&updatedUser); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        err := endpoints.UpdateUser(userID, updatedUser)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
    })

    router.DELETE("/User/:id", func(ctx *gin.Context) {
        
        userID := ctx.Param("id")
    
        err := endpoints.DeleteUser(userID)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
    })

    router.POST("/Vehicle", func(c *gin.Context) {
        var vehicle endpoints.Vehicles
        if err := c.BindJSON(&vehicle); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        id, err := endpoints.AddVehicle(vehicle)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"id": id})
    })

    router.GET("/Vehicle", func(ctx *gin.Context) {
        vehicles, err := endpoints.FetchVehicles()
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusOK, vehicles)
    })

    
    router.PUT("/Vehicle/:id", func(ctx *gin.Context) {
       
        vehicleID := ctx.Param("id")
    
        var updatedVehicle endpoints.Vehicles
        if err := ctx.BindJSON(&updatedVehicle); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        err := endpoints.UpdateVehicle(vehicleID, updatedVehicle)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Vehicle updated successfully"})
    })
    router.DELETE("/Vehicle/:id", func(ctx *gin.Context) {
        
        vehicleID := ctx.Param("id")
    
        err := endpoints.DeleteVehicle(vehicleID)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Vehicle deleted successfully"})
    })

    router.POST("/Client", func(c *gin.Context) {
        var client endpoints.Clients
        if err := c.BindJSON(&client); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        id, err := endpoints.AddClient(client)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"id": id})
    })

    router.GET("/Client", func(ctx *gin.Context) {
        clients, err := endpoints.FetchClients()
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusOK, clients)
    })

    
    router.PUT("/Client/:id", func(ctx *gin.Context) {
       
        clientID := ctx.Param("id")
    
        var updatedClient endpoints.Clients
        if err := ctx.BindJSON(&updatedClient); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        err := endpoints.UpdateClient(clientID, updatedClient)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Client updated successfully"})
    })
    router.DELETE("/Client/:id", func(ctx *gin.Context) {
        
        clientID := ctx.Param("id")
    
        err := endpoints.DeleteClient(clientID)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
    })


    router.POST("/Driver", func(c *gin.Context) {
        var driver endpoints.Drivers
        if err := c.BindJSON(&driver); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        id, err := endpoints.AddDriver(driver)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"id": id})
    })

    router.GET("/Driver", func(ctx *gin.Context) {
        drivers, err := endpoints.FetchDrivers()
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusOK, drivers)
    })

    
    router.PUT("/Driver/:id", func(ctx *gin.Context) {
       
        driverID := ctx.Param("id")
    
        var updatedDriver endpoints.Drivers
        if err := ctx.BindJSON(&updatedDriver); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        err := endpoints.UpdateDriver(driverID, updatedDriver)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Driver updated successfully"})
    })
    router.DELETE("/Driver/:id", func(ctx *gin.Context) {
        
        driverID := ctx.Param("id")
    
        err := endpoints.DeleteDriver(driverID)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
    })

    router.POST("/Trip", func(c *gin.Context) {
        var trip endpoints.Trips
        if err := c.BindJSON(&trip); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        id, err := endpoints.AddTrip(trip)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"id": id})
    })

    router.GET("/Trip", func(ctx *gin.Context) {
        trips, err := endpoints.FetchTrips()
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusOK, trips)
    })

    
    router.PUT("/Trip/:id", func(ctx *gin.Context) {
       
        tripID := ctx.Param("id")
    
        var updatedTrip endpoints.Trips
        if err := ctx.BindJSON(&updatedTrip); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        err := endpoints.UpdateTrip(tripID, updatedTrip)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Trip updated successfully"})
    })
    router.DELETE("/Trip/:id", func(ctx *gin.Context) {
        
        tripID := ctx.Param("id")
    
        err := endpoints.DeleteTrip(tripID)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Trip deleted successfully"})
    })

    router.POST("/Payment", func(c *gin.Context) {
        var payment endpoints.Payments
        if err := c.BindJSON(&payment); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        id, err := endpoints.AddPayment(payment)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"id": id})
    })

    router.GET("/Payment", func(ctx *gin.Context) {
        payments, err := endpoints.FetchPayment()
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusOK, payments)
    })

    
    router.PUT("/Payment/:id", func(ctx *gin.Context) {
       
        paymentID := ctx.Param("id")
    
        var updatedPayment endpoints.Payments
        if err := ctx.BindJSON(&updatedPayment); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        err := endpoints.UpdatePayment(paymentID, updatedPayment)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully"})
    })
    router.DELETE("/Payment/:id", func(ctx *gin.Context) {
        
        paymentID := ctx.Param("id")
    
        err := endpoints.DeletePayment(paymentID)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    
        ctx.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
    })

	router.Run("localhost:9030")
}