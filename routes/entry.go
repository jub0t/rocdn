package routes

import (
	rblx "rblx/api"
	"rblx/database"
	"rblx/structs"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	ValidSizes = []int16{30, 48, 60, 75, 100, 110, 140, 150, 352, 420, 720}
)

var (
  InMemoryCache = true; // If this is set to true the images will be stored in memory.
)


func PrimaryRoute(av *structs.Storage, hs *structs.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, structs.Response{
			Success: true,
			Message: "Welcome - Roblox Cache Server By github.com/jub0t/RoCDN",
			Data: structs.DatabaseInfo{
				Avatars:   len(av.Data),
				Headshots: len(hs.Data),
			},
		})
	}
}

func NotFound(c echo.Context) error {
	return c.Redirect(302, `https://github.com/jub0t/RoCDN`)
}

func Headshot(db *structs.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id, err := strconv.ParseInt(c.Param("userId"), 0, 64)

		if err != nil {
			return c.JSON(400, structs.Response{
				Success: true,
				Message: "Unable to parse user_id",
			})
		}

		size, err := strconv.ParseInt(c.QueryParam("size"), 0, 64)

		if err != nil {
			return c.JSON(400, structs.Response{
				Success: true,
				Message: "Invalid size",
			})
		}

		image := database.Get(db, int(user_id), int(size))

		if image.TargetId > 0 {
      if (InMemoryCache) {
        c.Response().Header().Set("Content-Type", "image/png")
        c.Response().Header().Set("Content-Length", strconv.Itoa(len(image.Data)))
        c.Response().Header().Set("Cache-Control", "max-age=31536000") // Cache for 1 year
        c.Response().Write(image.Data)
        c.Response().Flush()
        return nil;
      } else {
			  return c.Redirect(302, image.ImageUrl)
      }
		} else {
			r_image, err := rblx.GetHeadshot(int(user_id), int(size), "png", false)

      if err != nil {
				return c.JSON(400, structs.Response{
					Success: true,
					Message: "User not found",
				})
			}

      image_data, err := rblx.GetImageDataFromURL(r_image.ImageUrl);
			
      if err != nil {
				return c.JSON(400, structs.Response{
					Success: true,
					Message: "Image data could not be retreived",
				})
			}

      database.Insert(db, structs.Image{
        Data: image_data,
				Size:      int(size),
        TargetId:  r_image.TargetId,
				ImageUrl:  r_image.ImageUrl,
				Timestamp: time.Now().UnixMilli() + time.Hour.Milliseconds()*6,
			})

			if len(r_image.ImageUrl) > 0 {
        c.Response().Header().Set("Content-Type", "image/png")
        c.Response().Header().Set("Content-Length", strconv.Itoa(len(image_data)))
        c.Response().Header().Set("Cache-Control", "max-age=31536000") // Cache for 1 year
        c.Response().Write(image.Data)
        c.Response().Flush()
        return nil;
			} else {
				return c.JSON(400, structs.Response{
					Success: true,
					Message: "ImageURL Is Not Valid",
				})
			}
		}
	}
}

func Avatar(db *structs.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id, err := strconv.ParseInt(c.Param("userId"), 0, 64)

		if err != nil {
			return c.JSON(400, structs.Response{
				Success: true,
				Message: "Unable to parse user_id",
			})
		}

		size, err := strconv.ParseInt(c.QueryParam("size"), 0, 64)

		if err != nil {
			return c.JSON(400, structs.Response{
				Success: true,
				Message: "Invalid size",
			})
		}

		image := database.Get(db, int(user_id), int(size))

		if image.TargetId > 0 {
			return c.Redirect(302, image.ImageUrl)
		} else {
			r_image, err := rblx.GetAvatar(int(user_id), int(size), "png", false)

			if err != nil {
				return c.JSON(400, structs.Response{
					Success: true,
					Message: "User not found",
				})
			}

			database.Insert(db, structs.Image{
				Size:      int(size),
				TargetId:  r_image.TargetId,
				ImageUrl:  r_image.ImageUrl,
				Timestamp: time.Now().UnixMilli() + time.Hour.Milliseconds()*6,
			})

			if len(r_image.ImageUrl) > 0 {
				return c.Redirect(302, r_image.ImageUrl)
			} else {
				return c.JSON(400, structs.Response{
					Success: true,
					Message: "ImageURL Is Not Valid",
				})
			}
		}
	}
}
