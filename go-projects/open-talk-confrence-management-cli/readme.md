<h1 align=center>Open Talk Conference Management Cli</h1>
### Registration of the venue

1. Venue name
2. Venue place
3. Venue capacity
4. List of conference talks

### registration of the speaker

1. speaker name
2. venue and conference talk selection

### ticket booking for by the user

#### gathering user info

1. user first name
2. user second name
3. user email
4. selection of venue
5. selection of the booking slot
6. booking type

#### user validation

1. user first name and user last name should be greater than or equal to 2 character
2. check for email validation
3. check if the number of bookings is valid less than total capacity and remaining
4. venue selected should be from the list of venues
5. if any of the validation failed asked to enter that particular info only

#### Bookings Confirmation

1. Once the user validation is complete
2. show him the booking details with
   - user info
   - venue selected
   - number of booking
   - with the invoice details and booking price
   - ask for confirmation

#### Checkout And Saving The Booking

1. once the confirmation is done add the booking to confirmed booking list

#### sending emails for the bookings confirmed

1. create an async service to send the email to the user of the booking confirmation
2. the email template should contain the following info
   - user info
   - venue selected
   - number of booking
   - with the invoice details and booking price
