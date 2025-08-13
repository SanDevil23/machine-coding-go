# Parking Lot System

## Functions of the Parking Lot System

- Create the parking lot.
- Add floors to the parking lot.
- Add a parking lot slot to any of the floors.
- Given a vehicle:
  - Finds the first available slot.
  - Books the slot.
  - Creates a ticket.
  - Parks the vehicle.
  - Returns the ticket.
- Unparks a vehicle given the ticket ID.
- Display the number of free slots per floor for a specific vehicle type.
- Display all free slots per floor for a specific vehicle type.
- Display all occupied slots per floor for a specific vehicle type.

---

## Vehicle Details

- Every vehicle will have:
  - Type
  - Registration number
  - Color

### Vehicle Types

- Car
- Bike
- Truck

---

## Parking Slot Details

- Each slot type can park a specific type of vehicle.
- No other vehicle types are allowed in a slot.

### Finding the First Available Slot

- Slot must match the vehicle type.
- Slot should be on the **lowest possible floor**.
- Slot should have the **lowest possible slot number** on that floor.
- Slots are **numbered serially** from `1` to `n` on each floor.

---

## Parking Lot Floor Details

- Floors are **numbered serially** from `1` to `n`.
- Each floor may contain one or more parking slots of different types.
- **Slot allocation per floor**:
  - 1st slot: Truck
  - 2nd & 3rd slots: Bikes
  - Remaining slots: Cars

---

## Ticket Details

- Ticket ID format:  
